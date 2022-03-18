function getSessionID() {
    var url = window.location.href;
    var idx = url.lastIndexOf("=")
    var sessionID = url.substring(idx + 1);
    return sessionID;
}
