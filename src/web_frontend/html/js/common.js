var currentTeamID = null;
var currentTeamName = null;
var currentBattle = null;

function getSessionID() {
    var url = window.location.href;
    var idx = url.lastIndexOf("=")
    var sessionID = url.substring(idx + 1);
    return sessionID;
}

function getAuthHeader() {
    return {
        'Authorization': 'Bearer ' + getSessionID(),
    };
}

function getSessionStatus() {
    return new Promise((resolve, reject) => {
        $.ajax({
            url: `/api/v1/sessions/me`,
            headers: getAuthHeader(),
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

function challengeTeam(teamID) {
    return new Promise((resolve, reject) => {
        $.ajax({
            method: 'post',
            headers: getAuthHeader(),
            url: `/api/v1/challenges?team=${teamID}`,
            success: function () {
                resolve();
            },
            error: reject
        });
    });
}

function goToPageWithSession(page) {
    document.location.href = `/${page}.html?session=${getSessionID()}`;
}
