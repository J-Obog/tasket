import socketio

class Client:
    def __init__(self, url: str, type: str, channel: str, key: str = ""):
        headers = {"authorization": key, "client-type": type, "channel": channel}
        self._client = socketio.Client()
        self._client.connect(url, headers)

    def close(self):
        self._client.disconnect()
        