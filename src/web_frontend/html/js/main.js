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

$(function () {
    var currentTeamID = null;
    var $newTeamName = $('#teams input[name=name]');
    var $currentTeam = $('.current_team');
    var $confirmJoinTeam = $('#btnConfirmJoinTeam');
    var $joinTeamView = $('#teams .join');
    var $joinTeamsState = $joinTeamView.find('.state');
    var $joinTeamsList = $joinTeamView.find('ul');

    var $allViews = {
        'teamOverview': $('#teams .overview'),
        'newTeam': $('#teams .new'),
        'joinTeam': $joinTeamView,
        'lobby': $('#lobby'),
    };

    function switchView(view) {
        for (const v in $allViews) {
            if (v != view) {
                $allViews[v].hide();
            }
        }
        $allViews[view].find('.state').hide();
        $allViews[view].find('.loading').show();
        $allViews[view].show();
    }

    $('#btnCreateTeam').click(function () {
        $newTeamName.val('');
        switchView('newTeam');
    })

    $('#btnConfirmCreateTeam').click(function () {
        var teamName = $newTeamName.val();
        if (!teamName) {
            alert('Please enter a name for the team');
            return;
        }
        $.ajax({
            url: '/api/v1/teams',
            method: 'post',
            data: {
                name: teamName,
                owner: getSessionID(),
            },
            success: function (data) {
                switchToLobby(data.id, data.name)
            }
        });
    });

    $('#btnCancelCreateTeam, #btnCancelJoinTeam').click(function () {
        switchView('teamOverview');
    })

    $('#btnLeaveTeam').click(function () {
        alert('Implement btnLeaveTeam'); // XXX implement btn
    });

    $('#btnRefreshTeamsTable').click(function () {
        refreshTeamsStatus();
        refreshChallenges();
    });

    $('#btnJoinTeam').click(function () {
        $confirmJoinTeam.hide();
        switchView('joinTeam');
        $.ajax({
            url: '/api/v1/teams',
            success: function (teamList) {
                $joinTeamsState.hide();
                if (!teamList.length) {
                    $joinTeamView.find('.nonfound').show();
                } else {
                    $joinTeamsList.html('');
                    teamList.forEach(team => {
                        $(`<li><input type="radio" name="team" data-name="${team.name}" value="${team.id}"> ${team.name}</li>`).appendTo($joinTeamsList)
                    });
                    $joinTeamsList.show();
                    $confirmJoinTeam.show();
                }
            }
        })
    });

    $confirmJoinTeam.click(function () {
        var $selected = $joinTeamsList.find('input:checked');
        if (!$selected.length) {
            alert('Please select a team to join')
            return;
        }

        var teamID = $selected.val();
        var teamName = $selected.data('name');
        $.ajax({
            method: 'post',
            url: `/api/v1/teams/${teamID}/join`,
            headers: getAuthHeader(),
            success: function () {
                switchToLobby(teamID, teamName)
            }
        })
    });

    function switchToLobby(teamID, teamName) {
        currentTeamID = teamID;
        $currentTeam.text(teamName);
        switchView('lobby');
        refreshTeamsStatus();
        refreshChallenges();
    }

    function refreshTeamsStatus() {
        var $teamListTable = $('#team_list tbody');
        $.ajax({
            url: '/api/v1/teams',
            success: function (teamList) {
                $teamListTable.html('');
                teamList.forEach(team => {
                    // TODO: Use team.status.timestamp
                    var actionBtn = '';
                    if (team.id == currentTeamID) {
                        actionBtn = '(your team)';
                    } else if (team.status.battleID) {
                        actionBtn = `<button class=watchBattle data-battle-id=${team.status.battleID}>Watch battle</button>`;
                    } else {
                        actionBtn = `<button class=challengeTeam data-team-id=${team.id}>Challenge team</button>`;
                    }

                    $(`
                        <tr>
                            <td>${team.name}</td>
                            <td>${team.status.status}</td>
                            <td>${team.members}</td>
                            <td>${team.rank ? team.rank : '?'}</td>
                            <td>${actionBtn}</td>
                        </tr>
                    `).appendTo($teamListTable)
                });
            }
        })
    }

    function refreshChallenges() {
        var $challenges = $('#incoming-challenges');
        $challenges.hide();
        $.ajax({
            url: `/api/v1/challenges`,
            headers: getAuthHeader(),
            success: function (challenges) {
                if (!challenges.length) {
                    return;
                }
                $challenges.find('p span').text(challenges.length);
                var $list = $challenges.find('ul');
                $list.html('');
                challenges.forEach(c => {
                    $(`<li><button data-challenge-id="${c.id}">Accept</button> challenge from <span>${c.challenger.name}</span> (<span>${c.timestamp}</span> ago)</li>`).appendTo($list)
                });
                $challenges.show();
            }
        })
    }

    $('#incoming-challenges').on('click', 'button', function () {
        // XXX implement accept challenge button
        var challengeID = $(this).data('challenge-id');
        alert('TODO: Implement accept challenge - ID: ' + challengeID);
    })

    $('#team_list').on('click', '.watchBattle', function () {
        // XXX implement watch battle button
        var battleID = $(this).data('battle-id');
        alert('TODO: Implement watch battle - ID: ' + battleID);
    })

    $('#team_list').on('click', '.challengeTeam', function () {
        // XXX implement challenge team button
        var teamID = $(this).data('team-id');
        alert('TODO: Implement challenge team - ID: ' + teamID);
    })

});
