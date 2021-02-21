package com.example.ascii.Logic;

import java.io.IOException;

public class Handler {
    public void handle(int width,String readFrom,String saveTo,boolean inverse) throws IOException {
        Reader reader = new Reader();
        ToAscii toAscii = new ToAscii();
        Writer writer = new Writer();
        Photo photo = reader.loadPhoto(readFrom,1,true);
        long startTime = System.nanoTime();
        Photo[][] grid = photo.toChunks(width, (int) (photo.getProportionalLength(width)*0.5));
        StringBuilder result = new StringBuilder();
        for (int row=0;row<grid.length;row++){
            for (int col=0;col<grid[0].length;col++){
                result.append(toAscii.photoToAscii(grid[row][col],inverse));
            }
            result.append('\n');
        }
        long endTime = System.nanoTime();
        System.out.println(((endTime-startTime)/1000000)+"ms total run time.");
        writer.write(saveTo, result.toString());
    }

    public void handle(int width,String readFrom,String saveTo) throws IOException {
        Reader reader = new Reader();
        ToAscii toAscii = new ToAscii();
        Writer writer = new Writer();
        Photo photo = reader.loadPhoto(readFrom,1,false);
        long startTime = System.nanoTime();
        Photo[][] grid = photo.toChunks(width, (int) (photo.getProportionalLength(width)*0.5));
        StringBuilder result = new StringBuilder();
        for (int row=0;row<grid.length;row++){
            for (int col=0;col<grid[0].length;col++){
                result.append(toAscii.photoToAscii(grid[row][col],false));
            }
            result.append('\n');
        }
        long endTime = System.nanoTime();
        System.out.println(((endTime-startTime)/1000000)+"ms total run time.");
        writer.write(saveTo, result.toString());
    }
}
