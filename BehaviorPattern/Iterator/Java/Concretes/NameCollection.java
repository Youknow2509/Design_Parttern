package Concretes;

import Interfaces.*;

public class NameCollection implements Aggregate {
    private String[] names = {"Alice", "Bob", "Charlie", "Diana"};

    @Override
    public Iterator createIterator() {
        return new NameIterator();
    }

    private class NameIterator implements Iterator {
        private int index = 0;

        @Override
        public boolean hasNext() {
            return index < names.length;
        }

        @Override
        public Object next() {
            if (this.hasNext()) {
                return names[index++];
            }
            return null;
        }
    }
}

