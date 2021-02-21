package com.example.ascii.Logic;

import java.io.IOException;

public class Handler {
    public void handle(int width,String filename) throws IOException {
        Reader reader = new Reader();
        ToAscii toAscii = new ToAscii();
        Writer writer = new Writer();
        Photo photo = reader.loadPhoto("C:\\Users\\aarse\\Documents\\pngReader\\src\\eiffel.png",5);
        long startTime = System.nanoTime();
        Photo[][] grid = photo.toChunks(width, photo.getProportionalLength(width));
        StringBuilder result = new StringBuilder();
        for (int row=0;row<grid.length;row++){
            for (int col=0;col<grid[0].length;col++){
                result.append(toAscii.photoToAscii(grid[row][col]));
            }
            result.append('\n');
        }
        long endTime = System.nanoTime();
        System.out.println(((endTime-startTime)/1000000)+"ms total run time.");
        writer.write(filename, result.toString());
    }
}
