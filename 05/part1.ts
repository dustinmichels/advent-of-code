let STATE = {}

// ---------- LOAD DRAWING ----------

/**
 * Parse lines of drawing into array with len equal to number of stacks
 *   Eg, [ ".", ".", ".", "Q", ".", "G", ".", "M", "." ]
 */
function parseLine(line: string) {
  // replace empty block (3 spaces) with a dot
  const re = /(   ) /g
  line = line.replace(re, '.')

  // split into array of chars and dots
  const keepChars = /[A-Z.]/g
  return line.match(keepChars) as string[]
}

function addLineToState(lineArr: string[]) {
  for (let i = 0; i < lineArr.length; i++) {
    if (lineArr[i] !== '.') {
      STATE[i + 1].push(lineArr[i])
    }
  }
}

function initState(crateDrawing: string) {
  const lines = crateDrawing.split('\n')

  // init state using the ids line
  let idLine = lines.pop()
  const ids = intArrayFromString(idLine)
  STATE = Object.fromEntries(ids.map((x) => [x, []]))

  // Add the rows of the drawing to STATE
  while (lines.length > 0) {
    let l = lines.pop()
    let lineArr = parseLine(l)
    addLineToState(lineArr)
  }
}

// ----------- UPDATE STATE ----------

const move = (n_times: number, from_stack: number, to_stack: number) => {
  for (let i = 0; i < n_times; i++) {
    STATE[to_stack].push(STATE[from_stack].pop())
  }
}

const parseCommand = (inp: string): [number, number, number] => {
  const m = intArrayFromString(inp)
  return [m[0], m[1], m[2]]
}

function handleCommands(commandText: string) {
  const commands = commandText.split('\n')
  while (commands.length > 0) {
    const c = commands.shift()
    if (c == '') return
    move(...parseCommand(c))
  }
}

// ---------- HELPER FUNCTIONS ----------

const displayState = () => {
  console.log(STATE)
}

const intArrayFromString = (inp: string) => {
  const numberPattern = /\d+/g
  const m = inp.match(numberPattern)
  return m.map((x) => parseInt(x))
}

async function loadFile(filename: string) {
  const blob = Bun.file(filename)
  return await blob.text()
}

function getTopCrates() {
  let tops = ''
  for (let i = 0; i < Object.keys(STATE).length; i++) {
    let arr = STATE[i + 1]
    tops += arr[arr.length - 1]
  }
  return tops
}

// ---------- MAIN ----------

async function main() {
  const text = await loadFile('input.txt')
  const [crateDrawing, commandText] = text.split('\n\n')

  console.log('> init state')
  initState(crateDrawing)

  console.log('> handle commands')
  handleCommands(commandText)
  // displayState()

  console.log('> get top crates')
  const topCrates = getTopCrates()
  console.log(topCrates)
}

main()
