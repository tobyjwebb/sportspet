$(function () {
    var $board = $('#chessboard');
    drawChessboard($board);
    updateChessPieces($board, boardStatus);
});

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
            html += `<td class="${id} ${colour}"></td>`;
        }
        html += `</tr>`;
    }

    // letter row on bottom:
    html += `<tr><td></td>`
    for (var col = 0; col < 8; col++) {
        html += `
                <td>${letters[col]}</td>
            `;

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
