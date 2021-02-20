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
        if (chunkHeight==0){
            chunkHeight = 1;
        }
        if (chunkWidth==0){
            chunkWidth=1;
        }
        Photo[][] photoGrid = new Photo[imgData.length/chunkHeight+1][imgData[0].length/chunkWidth+1];
        for (int row=0;row< imgData.length;row+=chunkHeight){
            for (int col=0;col<imgData[0].length;col+=chunkWidth){
                double [][][]chunk = new double[chunkHeight][chunkWidth][4];
                for (int i=row;i<row+chunkHeight&&i<imgData.length;i++){
                    chunk[i-row]=Arrays.copyOfRange(imgData[i],col,col+chunkWidth);
                }
                photoGrid[row / chunkHeight][col / chunkWidth] = new Photo(chunk);
            }
        }
        return photoGrid;
    }
    
     public int getProportionalHeight(int length) {
        return length * this.imgData.length / this.imgData[0].length; 
    }
    
    public int getProportionalLength(int height) {
        return this.imgData[0].length * height / this.imgData.length;
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
