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

function addLineToState(state: object, lineArr: string[]) {
  for (let i = 0; i < lineArr.length; i++) {
    if (lineArr[i] !== '.') {
      state[i + 1].push(lineArr[i])
    }
  }
}

export function initState(crateDrawing: string) {
  const lines = crateDrawing.split('\n')

  // init state using the ids line
  let idLine = lines.pop()
  const ids = intArrayFromString(idLine)
  let state = Object.fromEntries(ids.map((x) => [x, []]))

  // Add the rows of the drawing to STATEs
  while (lines.length > 0) {
    let l = lines.pop()
    let lineArr = parseLine(l)
    addLineToState(state, lineArr)
  }

  return state
}

// ---------- HELPER FUNCTIONS ----------

const intArrayFromString = (inp: string) => {
  const numberPattern = /\d+/g
  const m = inp.match(numberPattern)
  return m.map((x) => parseInt(x))
}

export const parseCommand = (inp: string): [number, number, number] => {
  const m = intArrayFromString(inp)
  return [m[0], m[1], m[2]]
}

export async function loadFile(filename: string) {
  const blob = Bun.file(filename)
  return await blob.text()
}

export function getTopCrates(state: object) {
  let tops = ''
  for (let i = 0; i < Object.keys(state).length; i++) {
    let arr = state[i + 1]
    tops += arr[arr.length - 1]
  }
  return tops
}
