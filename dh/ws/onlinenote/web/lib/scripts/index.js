(function ($) {

    const VueApp = {
        data() {
            return {
                messages: ["hello", "world"],
                text: '',
                note: {
                    name: '',
                    content: ''
                },
                user: { id: 0, name: '' },
                errMessage: ''
            }
        },
        mounted() {
            this.init();
        },
        methods: {
            init() {
                let that = this;
                if (window["WebSocket"]) {
                    conn = new WebSocket("ws://" + document.location.host + "/ws");
                    console.log(document.location.host)
                    conn.onclose = function (evt) {
                        that.conn = null;
                        that.errMessage = "Connection closed.";
                    };
                    conn.onmessage = function (evt) {
                        that.onmessage(evt.data);
                    };
                    that.conn = conn;
                } else {
                    that.errMessage = "Your browser does not support web socket.";
                }
            },
            onmessage(data) {
                this.messages.push(data);
            },
            chatKeyDown(event){
                if(event.key==='Enter'){
                    this.sendChat();
                }
            },
            sendChat(){
                console.log("sendChat", this.text);
                this.send(this.text);
                this.text = '';
            },
            send(msg) {
                let that = this;
                if (!that.conn) {
                    that.errMessage = "Connection not ready.";
                    return false;
                }
                this.conn.send(msg);
                return true;
            },
            hello: function(){
                console.log("hello");
            }
        },
    };

    Vue.createApp(VueApp).mount('#vue-main');

})(jQuery)