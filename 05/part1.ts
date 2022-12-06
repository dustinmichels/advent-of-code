import { initState, loadFile, parseCommand, getTopCrates } from './base'

let STATE = {}

// const displayState = () => {
//   console.log(STATE)
// }

// ----------- UPDATE STATE ----------

const move = (n_times: number, from_stack: number, to_stack: number) => {
  for (let i = 0; i < n_times; i++) {
    STATE[to_stack].push(STATE[from_stack].pop())
  }
}

export function handleCommands(commandText: string) {
  const commands = commandText.split('\n')
  while (commands.length > 0) {
    const c = commands.shift()
    if (c == '') return
    move(...parseCommand(c))
  }
}

// ---------- MAIN ----------

async function main() {
  const text = await loadFile('input.txt')
  const [crateDrawing, commandText] = text.split('\n\n')

  console.log('> init state')
  STATE = initState(crateDrawing)

  console.log('> handle commands')
  handleCommands(commandText)
  // displayState()

  console.log('> get top crates')
  const topCrates = getTopCrates(STATE)
  console.log(topCrates)
}

main()
