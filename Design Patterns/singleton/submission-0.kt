class Singleton private constructor() {
    private var value: String? = null;
    companion object {
        private var instance: Singleton = Singleton()

        fun getInstance(): Singleton {
            return instance
        }
    }

    fun getValue(): String? {
        return value
    }

    fun setValue(value: String) {
        this.value = value
    }
}
