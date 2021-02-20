import java.util.ArrayList;
import java.util.Arrays;

public class Photo {
    private double[][][] imgData;
    public Photo(double[][][] imgData){
        this.imgData = imgData;
    }

    public double[][][] getImgData() {
        return imgData;
    }

    public Photo[][] toChunks (int width,int height){
        int chunkHeight = imgData.length/height;
        int chunkWidth = imgData[0].length/width;
        Photo[][] photoGrid = new Photo[imgData.length/chunkHeight+1][imgData[0].length/chunkWidth+1];
        for (int row=0;row< imgData.length;row+=chunkHeight){
            for (int col=0;col<imgData[0].length;col+=chunkWidth){
                double[][][] chunk = new double[chunkHeight][chunkWidth][4];
                for (int subRow=row;subRow<row+chunkHeight&&subRow<imgData.length;subRow++){
                    double[][] chunkLine = new double[chunkWidth][4];
                    for (int subCol=col;subCol<col+chunkWidth&&subCol<imgData[0].length;subCol++){
                        chunkLine[subCol-col]=imgData[subRow][subCol];
                    }
                    chunk[subRow-row] = chunkLine;
                }
                photoGrid[row/chunkHeight][col/chunkWidth] = new Photo(chunk);
            }
        }
        return photoGrid;
    }

    public void prettyPrint(){
        for (int i=0;i<imgData.length;i++){
            for (int b=0;b<imgData[0].length;b++){
                System.out.println(Arrays.toString(imgData[i][b]));
            }
        }
    }

    public void printInfo(){
        System.out.println(imgData.length);
        System.out.println(imgData[0].length);
    }
}
