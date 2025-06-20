/**
 * 分片上传
 * 2025.06.17 by dralee
 */
const CHUNK_SIZE = 2 * 1024 * 1024; // 2MB
const MAX_CONCURRENCY = 3;

function log(msg) {
    const logBox = document.getElementById("log");
    logBox.innerHTML += msg + "<br>";
    logBox.scrollTop = logBox.scrollHeight;
}

function clear(){
    const logBox = document.getElementById("log");
    logBox.innerHTML = "";
}

function getSha1Hash(input) {
    return CryptoJS.SHA1(input).toString(CryptoJS.enc.Hex);
}

function updateProgress(percent) {
    document.getElementById("bar").style.width = percent + "%";
}

function sliceFile(file) {
    const chunks = [];
    const totalChunks = Math.ceil(file.size / CHUNK_SIZE);
    for (let i = 0; i < totalChunks; i++) {
        const start = i * CHUNK_SIZE;
        const end = Math.min(file.size, start + CHUNK_SIZE);
        const blob = file.slice(start, end);
        const chunkName = getSha1Hash(file.name + '-' + i);
        chunks.push({ index: i, blob, chunkName });
    }
    return chunks;
}

async function uploadChunk(chunk, file, fileID) {
    const formData = new FormData();
    // formData.append('file', chunk.blob);
    // formData.append('filename', file.name);
    // formData.append('chunkIndex', chunk.index);
    // formData.append('chunkName', chunk.chunkName);

    formData.append("file_id", fileID);
    formData.append("index", chunk.index);
    formData.append("chunk", chunk.blob);

    const res = await fetch('/upload', {
        method: 'POST',
        body: formData
    });

    if (!res.ok) throw new Error(`Chunk ${chunk.index} 上传失败`);
    log(`✔ 上传分片 ${chunk.index} 成功`);
}

async function checkChunk(fileId, chunkIndex) {
    const res = await fetch(`/checkchunk?file_id=${fileId}&index=${chunkIndex}`);
    return res.ok;
}


async function uploadChunksConcurrently(chunks, file, fileId) {
    let index = 0;
    let count = 0;
    async function worker() {
        while (index < chunks.length) {
            const current = chunks[index++];
            try {
                if(await checkChunk(fileId, current.index)) {
                    log(`✔ 分片 ${current.index} 已上传，跳过`);
                }else{
                    await uploadChunk(current, file, fileId);
                }
                count++;
                console.log("count:", count, "total: ", chunks.length);
                updateProgress(count/chunks.length*100);
            } catch (err) {
                log(`❌ 分片 ${current.index} 失败: ${err.message}`);
                throw err;
            }
        }
    }

    const workers = [];
    for (let i = 0; i < MAX_CONCURRENCY; i++) {
        workers.push(worker());
    }
    await Promise.all(workers);
}

async function startUpload() {
    const file = document.getElementById('fileInput').files[0];
    if (!file) return alert('请选择文件');

    clear();
    log(`📦 开始上传 ${file.name}，大小 ${(file.size / 1024 / 1024).toFixed(2)} MB`);

    const chunks = sliceFile(file);
    log(`🔪 分片数：${chunks.length}`);

    const fileId = getSha1Hash(file.name + file.size);
    log(`🔑 文件 ID：${fileId}`);

    try {
    await uploadChunksConcurrently(chunks, file, fileId);
    log("📡 分片全部上传完成，通知服务器合并...");

    const mergeForm = new FormData();
    mergeForm.append("file_id", fileId);
    mergeForm.append("total_chunks", chunks.length);
    mergeForm.append("filename", file.name);

    log(`🔪 开始合并分片...${fileId}, ${chunks.length},${file.name}`);

    const res = await fetch("/merge", {
        method: "POST",
        body: mergeForm
    });
    if (res.ok) {
        log("🎉 合并成功！");
    } else {
        log("❌ 合并失败！");
    }
    } catch (err) {
    log("❌ 上传过程中出现错误：" + err.message);
    }
}