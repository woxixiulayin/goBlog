export default function assert(condition, errorMessage = "condition") {
    if (!condition) {
        throw new Error(`[assert] ${errorMessage}`)
    }
}
