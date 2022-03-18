var currentTeamID = null;
var currentTeamName = null;
var currentBattle = null;

function getSessionID() {
    var url = window.location.href;
    var idx = url.lastIndexOf("=")
    var sessionID = url.substring(idx + 1);
    return sessionID;
}

function getSessionStatus() {
    var sessionID = getSessionID();
    return new Promise((resolve, reject) => {
        $.ajax({
            url: `/api/v1/sessions/${sessionID}`,
            success: function (res) {
                currentBattle = res.battle;
                if (res.team) {
                    currentTeamID = res.team.id;
                    currentTeamName = res.team.name;
                } else {
                    currentTeamID = null;
                    currentTeamName = null;
                }
                resolve();
            },
            error: reject
        });
    });
}


function getBattleStatus(battleID) {
    return new Promise((resolve, reject) => {
        $.ajax({
            url: `/api/v1/battles/${battleID}/state`,
            success: function (res) {
                resolve(res);
            },
            error: reject
        });
    });
}

function sendBattleMovement(battleID, from, to) {
    return new Promise((resolve, reject) => {
        $.ajax({
            method: 'post',
            url: `/api/v1/battles/${battleID}/move`,
            data: { from, to },
            success: function (res) {
                resolve(res);
            },
            error: reject
        });
    });
}
