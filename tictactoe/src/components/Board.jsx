import Square from "./Square"
import styles from "./Board.module.css"

/*
 * Board component: Renders the Tic-Tac-Toe board.
 * Receives the `squares` array (representing the board state) and the `onClick` function
 * from the parent component (`App`) to handle square clicks.
 */

const Board = ({ squares, onClick }) => {
  const renderSquare = (i) => <Square value={squares[i]} onClick={() => onClick(i)} />

  return (
    <div className={styles.board}>
      {[0, 1, 2].map((row) => (
        <div key={row} className={styles.boardRow}>
          {[0, 1, 2].map((col) => renderSquare(row * 3 + col))}
        </div>
      ))}
    </div>
  )
}

export default Board