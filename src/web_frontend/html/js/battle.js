var currentTeamBlack = false;
var currentTeamWhite = false;
let current_team_turn = false;
$(function () {
    var $board = $('#chessboard');
    // We need to get the current team colour before drawing the board,
    // to see if we need to invert it.
    getSessionStatus().then(() => {
        let inverted = false;
        getBattleStatus(currentBattle).then(res => {
            if (currentTeamID == res.black_team) {
                inverted = true;
                currentTeamBlack = true;
            } else if (currentTeamID == res.white_team) {
                currentTeamWhite = true;
            }
            drawChessboard($board, inverted);
            refreshBoardStatusPeriodic();
        });
    });

    $board.on('click', 'td.black, td.white', function () {
        var $this = $(this);
        let piece = $this.text();
        if (!current_team_turn) {
            return;
        }
        // If starting movement, and cannot move piece, do nothing.
        if (!movement && !canMovePiece(piece)) {
            return;
        }
        var cellID = $this.data('cell');
        // alert(`Clicked on cell: ${cellID}`);
        $this.toggleClass('selected');
        if (movement) {
            if (movement.from == cellID) {
                movement = null;
            } else {
                movement.to = cellID;
                console.log(`Sending move: ${movement.from}->${movement.to}`);
                sendBattleMovement(currentBattle, movement.from, movement.to).then(refreshBoardStatus);
                movement = null;
                $board.find('.selected').removeClass('selected');
            }
        } else {
            movement = { from: cellID };
        }
    })

    $('#btnLeaveBattle').click(function () {
        goToPageWithSession('main');
    });
});

var movement = null; // {from,to}

var letters = 'abcdefgh';
var pieceSprites = {
    p: '♟︎', // Black pawn
    c: '♜', // castle
    q: '♛', // queen
    b: '♝', // bishop
    h: '♞', // horse
    k: '♚', // king
    P: '♙', // White pawn
    C: '♖', // ...
    Q: '♕',
    B: '♗',
    H: '♘',
    K: '♔',
};

var boardStatus = "                                                                ";

function drawChessboard($board, inverted) {
    var html = '';

    for (var row = 8; row > 0; row--) {
        // TODO: re-add the coordinates
        // html += `<tr><td>${row}</td>`;
        html += `<tr>`;
        for (var col = 1; col <= 8; col++) {
            // We invert the rows/columns if the board is inverted, so the pieces are drawn in the correct coordinates.
            let id = inverted ?
                `${letters[8 - col]}${9 - row}` :
                `${letters[col - 1]}${row}`;
            let colour = (row + col) % 2 ? 'white' : 'black';
            html += `<td data-cell="${id}" class="${id} ${colour}"></td>`;
        }
        html += `</tr>`;
    }

    // Letter row on bottom:
    // html += `<tr><td></td>`
    // for (var col = 0; col < 8; col++) {
    //     html += `<td>${letters[col]}</td>`;
    // }
    // html += `</tr>`

    $board.html(`<tbody>${html}</tbody>`)
}

function updateChessPieces($board, pieces) {
    if (pieces.length != 64) {
        alert('Cant redraw chess pieces - wrong board size!');
        return;
    }

    let offset = 0;
    for (var row = 1; row <= 8; row++) {
        for (var col = 0; col < 8; col++) {
            let cell = letters[col] + row;
            let piece = pieces[offset];
            if (pieceSprites[piece]) {
                piece = pieceSprites[piece];
            }
            $board.find(`.${cell}`).text(piece);
            offset++;
        }
    }
}

function refreshBoardStatusPeriodic() {
    refreshBoardStatus();
    setInterval(refreshBoardStatus, 1500);
}

function refreshBoardStatus() {
    getBattleStatus(currentBattle).then(res => {
        updateChessPieces($('#chessboard'), res.board);
        updateTurnStatus(res);
    })
}

function updateTurnStatus(boardStatus) {
    var $battle_status = $('#battle_status');
    if (boardStatus.black_team == currentTeamID && boardStatus.turn == "black") {
        $battle_status.text("Your team's turn!");
        current_team_turn = true;
    } else if (boardStatus.white_team == currentTeamID && boardStatus.turn == "white") {
        $battle_status.text("Your team's turn!");
        current_team_turn = true;
    } else {
        $battle_status.text(`Team ${boardStatus.turn}'s turn.`);
        current_team_turn = false;
    }
    if (current_team_turn) {
        $('#chessboard').addClass('your-turn');
    } else {
        $('#chessboard').removeClass('your-turn');
    }
}

function sprite2letter(sprite) {
    for (let letter in pieceSprites) {
        if (pieceSprites[letter] == sprite) {
            return letter;
        }
    }
    return ' ';
}

function canMovePiece(p) {
    // p is a sprite, let's check the letter it represents to check if it is a lower or upper case:
    p = sprite2letter(p);
    if (p == ' ') {
        return false;
    }
    if (currentTeamWhite && p.toUpperCase() == p) {
        return true;
    }
    if (currentTeamBlack && p.toLowerCase() == p) {
        return true;
    }
    return false;
}
