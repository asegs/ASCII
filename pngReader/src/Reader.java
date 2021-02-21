import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.awt.image.Raster;
import java.io.File;
import java.io.IOException;
import java.util.Arrays;

public class Reader {
    public Photo loadPhoto(String filename,double contrastBoost) throws IOException {
        String prefix = filename.substring(filename.length() - 4, filename.length()); 
        if(prefix.equals(".png")){
            filename = jpegToPng(filename, filename.substring(0, filename.length() - 4).concat(".png"));
        }

        Utils utils = new Utils();
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
    
     public String jpegToPng(String inputFile, String outputFile) {
       try {
            BufferedImage bufferedImage = ImageIO.read(new File(inputFile));
            ImageIO.write(bufferedImage, "png", new File(outputFile));
        } catch (IOException e) {
            System.out.println(e);
        }
        return outputFile;
    }
}
}
