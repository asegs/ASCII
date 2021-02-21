package com.example.ascii.Logic;

import java.io.FileWriter;
import java.io.IOException;

public class Writer {
    public void write(String filename,String body){
        try {
            FileWriter myWriter = new FileWriter(filename);
            myWriter.write(body);
            myWriter.close();
            System.out.println("Successfully wrote to the file.");
        } catch (IOException e) {
            System.out.println("Filename: "+filename);
            System.out.println("File body: "+body);
            System.out.println("An error occurred.");
            e.printStackTrace();
        }
    }
}
