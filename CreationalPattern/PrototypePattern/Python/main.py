"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Prototype Pattern
"""

class Person:
    """
        Person class
    """
    def __init__(self, *args, **kwargs):
        self.name = kwargs.get('name','None')
        self.age = kwargs.get('age','None')
        self.gender = kwargs.get('gender','None')
        self.birthday = kwargs.get('birthday','None')
        self.contact = kwargs.get('contact','None')
        self.phone = kwargs.get('phone','None')
        self.email = kwargs.get('email','None')

        if args:
            lenArgs = len(args)
            if lenArgs >= 1:
                self.name = args[0]
            if lenArgs >= 2:
                self.age = args[1]
            if lenArgs >= 3:
                self.gender = args[2]
            if lenArgs >= 4:
                self.birthday = args[3]
            # TODO compelte it

    def clone(self):
        """
            Clone method
        """
        return Person(self)

    def __str__(self):
        return f'Name: {self.name}, Age: {self.age}, Phone: {self.phone}, Email: {self.email}'

    # Getter and setter methods
    def get_name(self):
        return self.name
    def set_name(self, name):
        self.name = name
    def get_age(self):
        return self.age
    def set_age(self, age):
        self.age = age
    def get_phone(self):
        return self.phone
    def set_phone(self, phone):
        self.phone = phone
    def get_email(self):
        return self.email
    def set_email(self, email):
        self.email = email
    # TODO complete it

def main():
    """
        Main function
    """

    person1 = Person(name='Ly Tran Vinh', age=20, phone='0123456789', email='lytranvinh.work@gmai.com')
    print(person1)

    person2 = person1.clone()
    person2.set_email('abc@gmail.com')
    person2.set_phone('0987654321')
    person2.set_name('Vinh')
    print(person2)

if __name__ == '__main__':
    main()