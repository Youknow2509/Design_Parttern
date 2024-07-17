"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Composite Pattern
"""

"""
    Deloy with manager File and Folder
"""
from Interface import File, Folder

def main():
    """
        Main function
    """
    # Create file
    file1 = File("File 1")
    file2 = File("File 2")
    file3 = File("File 3")
    
    # Create folder
    folder1 = Folder("Folder 1")
    folder2 = Folder("Folder 2")
    
    # Add file to folder
    folder1.addFile(file1)
    folder1.addFile(file2)
    folder2.addFile(file3)
    
    # Add folder to folder
    folder1.addFile(folder2)
    
    # Show all the details
    folder1.getDetails()
    
    # Remove file from folder
    folder1.removeFile(file2)
    
    # Show all the details
    folder1.getDetails()
    
if __name__ == '__main__':
    main()