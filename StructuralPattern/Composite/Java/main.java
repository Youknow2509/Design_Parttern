/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Bridge Pattern
 */

/**
 * Bridge Pattern with Folder and File
 */

import Interface.FileSystemComponent;
import Concrete.Folder;
import Concrete.File;

public class main {
    public static void main(String[] args) {
        // Creating files
        File file1 = new File("File1.txt", "txt", 12.0);
        File file2 = new File("File2.txt", "txt", 15.0);
        File file3 = new File("File3.txt", "txt", 16.0);

        // Creating folders
        Folder folder1 = new Folder("Folder1");
        Folder folder2 = new Folder("Folder2");
        Folder folder3 = new Folder("Folder3");

        // Composing the file system
        folder1.add(file1);
        folder1.add(file2);

        folder2.add(file3);
        folder2.add(folder1);

        folder3.add(folder2);

        // Displaying the file system structure
        folder3.showDetails();
    }
}