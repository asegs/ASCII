package com.example.ascii.Logic;
import java.awt.image.BufferedImage;
import java.io.File;
import javax.imageio.ImageIO;
import java.io.ByteArrayOutputStream;
import java.io.IOException;

public class Utils {
    public double boostContrastHex(double original,double boost,double average){
        double diff = original-average;
        double absDiff = Math.abs(diff);
        int sign = (int) ((diff)/absDiff);
        double result = original + (absDiff*sign*boost);
        if (result>255){
            result = 255;
        }else if (result<0){
            result = 0;
        }
        return result;
    }

    public String toPNG(String inputFile,String outputFile) throws IOException {
        System.out.println(inputFile);
        System.out.println(outputFile);
        BufferedImage bufferedImage = ImageIO.read(new File(inputFile));
        ImageIO.write(bufferedImage, "png", new File(outputFile));
        return outputFile;
    }

    public String PNGtoJPG(String inputFile,String outputFile) throws IOException {
        System.out.println(inputFile);
        System.out.println(outputFile);
        BufferedImage bufferedImage = ImageIO.read(new File(inputFile));
        ImageIO.write(bufferedImage, "jpg", new File(outputFile));
        System.out.println("WRITTEN");
        return outputFile;
    }

}
