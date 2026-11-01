class Solution {

    fun encode(strs: List<String>): String {
        return buildString {
            strs.forEach{
                append(it.length)
                append('#')
                append(it)
            }
        }
    }

    fun decode(str: String): List<String> {
        val result = mutableListOf<String>()
        var cur = 0
        while (cur < str.length) {
            var numString = ""
            while (str[cur].isDigit()) {
                numString += str[cur]
                cur++
            }
            if (str[cur] != '#') error("wrong format")
            cur++
            val len = numString.toInt()
            result.add(str.slice(cur..<cur + len))
            cur += len
        }
        return result
    }
}
