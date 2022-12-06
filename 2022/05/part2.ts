import { initState, loadFile, parseCommand, getTopCrates } from './base'

let STATE: Record<number, string[]> = {}

// For debug
const displayState = () => {
  console.log(STATE)
}

// ----------- UPDATE STATE ----------

const move = (n: number, fromStackRef: number, toStackRef: number) => {
  let fromStack = STATE[fromStackRef]
  let toStack = STATE[toStackRef]

  toStack = toStack.concat(fromStack.slice(fromStack.length - n))
  fromStack = fromStack.slice(0, fromStack.length - n)

  STATE[fromStackRef] = fromStack
  STATE[toStackRef] = toStack
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
  // const text = await loadFile('data/example_input.txt')
  const text = await loadFile('data/input.txt')

  const [crateDrawing, commandText] = text.split('\n\n')

  console.log('> init state')
  STATE = initState(crateDrawing)

  console.log('> handle commands')
  handleCommands(commandText)

  console.log('> get top crates')
  const topCrates = getTopCrates(STATE)
  console.log(topCrates)
}

main()
