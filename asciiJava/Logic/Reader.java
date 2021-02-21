package com.example.ascii.Logic;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.awt.image.Raster;
import java.io.File;
import java.io.IOException;

public class Reader {
    public Photo loadPhoto(String filename, double contrastBoost,boolean fixPNG) throws IOException {

        Utils utils = new Utils();
        final String noExtension = filename.substring(0, filename.length() - 4);
        String extension = filename.substring(filename.length()-4);
        if (extension.equals(".jpg")){
            utils.toPNG(filename,noExtension+".png");
        }
        if (fixPNG) {
            utils.PNGtoJPG(filename, noExtension + ".jpg");
            utils.toPNG(noExtension + ".jpg", filename);
        }
        BufferedImage image = ImageIO.read(new File(filename));
        Raster raster = image.getRaster();
        double[][][] imgData = new double[raster.getHeight()][raster.getWidth()][4];
        for (int row = 0; row < raster.getHeight(); row++) {
            for (int col = 0; col < raster.getWidth(); col++) {
                raster.getPixel(col, row, imgData[row][col]);
                double average = (imgData[row][col][0] + imgData[row][col][1] + imgData[row][col][2]) / 3;
                for (int i=0;i<3;i++){
                    imgData[row][col][i] = utils.boostContrastHex(imgData[row][col][i],contrastBoost,average);
                }
            }
        }
        return new Photo(imgData);
    }
}
