import { useSpring, animated, config } from "react-spring"
import styles from "./Square.module.css"

/*
 * Square component: Represents a single square on the Tic-Tac-Toe board.
 * It receives the `value` (X or O) and the `onClick` function as props.
 */

const Square = ({ value, onClick }) => {
  const animation = useSpring({
    from: { opacity: 0, transform: "scale(0.5) rotate(-180deg)" },
    to: { opacity: 1, transform: "scale(1) rotate(0deg)" },
    reset: true,
    config: config.wobbly,
  })

  return (
    <button className={styles.square} onClick={onClick}>
      <animated.span style={animation} className={styles.squareContent}>
        {value}
      </animated.span>
    </button>
  )
}

export default Square