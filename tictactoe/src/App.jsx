import { useState } from "react"
import Board from "./components/Board"
import { calculateWinner } from "./utils/GameLogic"
import styles from "./App.module.css"
import { useSpring, animated } from "react-spring"

const App = () => {
  const [history, setHistory] = useState([Array(9).fill(null)]) // For history of movies
  const [stepNumber, setStepNumber] = useState(0) // Current step number
  const [xIsNext, setXIsNext] = useState(true) // Is X the next player or O is the next player

  const handleClick = (i) => {
    const newHistory = history.slice(0, stepNumber + 1) // Keep history up to current step
    const current = newHistory[newHistory.length - 1] // Current board state
    const squares = [...current] // Copy of current board state

    // If there is a winner or the square is already filled, do nothing

    if (calculateWinner(squares) || squares[i]) {
      return
    }

    squares[i] = xIsNext ? "X" : "O"  // Update the square
    setHistory([...newHistory, squares])  // Update the history by appending the squares array to history
    setStepNumber(newHistory.length)  // Move to the next step
    setXIsNext(!xIsNext) // Switch players
  }

  const jumpTo = (step) => {
    setStepNumber(step)
    setXIsNext(step % 2 === 0) // Set correct player turn based on step
  }

  const resetGame = () => {
    setHistory([Array(9).fill(null)])
    setStepNumber(0)
    setXIsNext(true)
  }

  const current = history[stepNumber] // Current board state
  const winner = calculateWinner(current) // Check for a winner

  // Create move history list items, the buttons to which you can move to a state
  const moves = history.map((_, move) => (
    <li key={move}>
      <button className={styles.historyButton} onClick={() => jumpTo(move)}>
        {move ? `Move #${move}` : "Start"}
      </button>
    </li>
  ))

  // Determine game status message
  let status
  if (winner) {
    status = `Winner: ${winner}`
  } else if (stepNumber === 9) {
    status = "Draw!"
  } else {
    status = `Next player: ${xIsNext ? "X" : "O"}`
  }

  // React Spring animation for fade-in effect
  const fadeIn = useSpring({
    from: { opacity: 0 },
    to: { opacity: 1 },
    config: { duration: 1000 },
  })

  return (
    <animated.div style={fadeIn} className={styles.game}>
      <div className={styles.gameContainer}>
        <div className={styles.gameBoard}>
          <Board squares={current} onClick={handleClick} />
        </div>
        <div className={styles.gameInfo}>
          <div className={styles.status}>{status}</div>
          <button className={styles.resetButton} onClick={resetGame}>
            Reset Game
          </button>
          <ol className={styles.history}>{moves}</ol>
        </div>
      </div>
    </animated.div>
  )
}

export default App