$(function () {
    var $board = $('#chessboard');
    drawChessboard($board);
    updateChessPieces($board, boardStatus);
    getSessionStatus().then(() => {
        refreshBoardStatusPeriodic();
    })

    $board.on('click', 'td.black, td.white', function () {
        var $this = $(this);
        if (!movement && $this.text() == ' ') {
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

var boardStatus = "CHBQKBHCPPPPPPPP                                ppppppppchbqkbhc";

function drawChessboard($board) {
    var html = '';

    for (var row = 8; row > 0; row--) {
        html += `<tr><td>${row}</td>`;
        for (var col = 1; col <= 8; col++) {
            let id = `${letters[col - 1]}${row}`;
            let colour = (row + col) % 2 ? 'white' : 'black';
            html += `<td data-cell="${id}" class="${id} ${colour}"></td>`;
        }
        html += `</tr>`;
    }

    // Letter row on bottom:
    html += `<tr><td></td>`
    for (var col = 0; col < 8; col++) {
        html += `<td>${letters[col]}</td>`;
    }
    html += `</tr>`

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
    let your_turn = true;
    var $battle_status = $('#battle_status');
    if (boardStatus.black_team == currentTeamID && boardStatus.turn == "black") {
        $battle_status.text("Your team's turn!");
    } else if (boardStatus.white_team == currentTeamID && boardStatus.turn == "white") {
        $battle_status.text("Your team's turn!");
    } else {
        $battle_status.text(`Team ${boardStatus.turn}'s turn.`);
        your_turn = false;
    }
    if (your_turn) {
        $('#chessboard').addClass('your-turn');
    } else {
        $('#chessboard').removeClass('your-turn');
    }
}
