class SupportRequest:
    def __init__(self, message, complexity):
        self.message = message
        self.complexity = complexity

    def get_message(self):
        return self.message

    def get_complexity(self):
        return self.complexity
