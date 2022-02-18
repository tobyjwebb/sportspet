$(function () {
    var currentTeamID = null;
    var $newTeamName = $('#teams input[name=name]');
    var $currentTeam = $('#current_team');
    var $confirmJoinTeam = $('#btnConfirmJoinTeam');
    var $joinTeamView = $('#teams .join');
    var $joinTeamsState = $joinTeamView.find('.state');
    var $joinTeamsList = $joinTeamView.find('ul');

    var $allViews = {
        'teamOverview': $('#teams .overview'),
        'newTeam': $('#teams .new'),
        'joinTeam': $joinTeamView,
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
            data: { name: teamName },
            success: function (data) {
                $currentTeam.text(data.name);
                currentTeamID = data.id;
                switchView('teamOverview');
            }
        });
    });

    $('#btnCancelCreateTeam, #btnCancelJoinTeam').click(function () {
        switchView('teamOverview');
    })

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


        alert('XXX implement joinTeam()')
        // currentTeamID = $selected.val();
        // $currentTeam.text($selected.data('name'));

        // switchView('teamOverview');
    })

});
