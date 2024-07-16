"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Builder Pattern
"""

class Person:
    def __init__(self):
        self.name = None
        self.email = None
        self.phone = None
        self.age = None
        self.favoriteFilm = None
        self.favoriteSong = None
        self.favoriteBook = None

    def setName(self, name):
        self.name = name
        return self
    
    def setEmail(self, email):
        self.email = email
        return self
    
    def setPhone(self, phone):
        self.phone = phone
        return self
    
    def setAge(self, age):
        self.age = age
        return self
    
    def setFavoriteFilm(self, favoriteFilm):
        self.favoriteFilm = favoriteFilm
        return self
    
    def setFavoriteSong(self, favoriteSong):
        self.favoriteSong = favoriteSong
        return self
    
    def setFavoriteBook(self, favoriteBook):
        self.favoriteBook = favoriteBook
        return self

    def build(self):
        return self

    def __str__(self):
        return (f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}, Age: {self.age}, "
                f"Favorite Film: {self.favoriteFilm}, Favorite Song: {self.favoriteSong}, Favorite Book: {self.favoriteBook}")

def main():
    """
    Function main 
    """
    p = (Person()
         .setName("Ly Tran Vinh")
         .setEmail("lytranvinh.work@gmail.com")
         .setPhone("0123456789")
         .setAge(22)
         .setFavoriteFilm("The Shawshank Redemption")
         .setFavoriteSong("Shape of You")
         .setFavoriteBook("The Alchemist")
         .build())

    print(p)

if __name__ == "__main__":
    main()
