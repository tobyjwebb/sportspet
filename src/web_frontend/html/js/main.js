$(function () {
    var $newTeamName = $('#teams input[name=name]');
    var $currentTeam = $('#current_team');
    var $confirmJoinTeam = $('#btnConfirmJoinTeam');
    var $allViews = {
        'teamOverview': $('#teams .overview'),
        'newTeam': $('#teams .new'),
        'joinTeam': $('#teams .join'),
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
        alert(`XXX implement POST createTeam (name=${teamName})`);
        $currentTeam.text(teamName);
        switchView('teamOverview');
    });

    $('#btnCancelCreateTeam, #btnCancelJoinTeam').click(function () {
        switchView('teamOverview');
    })

    $('#btnJoinTeam').click(function () {
        $confirmJoinTeam.hide();
        switchView('joinTeam');
    });

});
