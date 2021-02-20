import java.io.IOException;
import java.util.Random;

public class Main {
    public static void main(String[] args) throws IOException {
        Reader reader = new Reader();
        ToAscii toAscii = new ToAscii();
        Writer writer = new Writer();
        Photo photo = reader.loadPhoto("C:\\Users\\aarse\\Documents\\pngReader\\src\\slight6.png",1.5);
        long startTime = System.nanoTime();
        Photo[][] grid = photo.toChunks(600, 180);
        String result = "";
        for (int row=0;row<grid.length;row++){
            for (int col=0;col<grid[0].length;col++){
                result+=toAscii.photoToAscii(grid[row][col]);
            }
            result+='\n';
        }
        long endTime = System.nanoTime();
        System.out.println(((endTime-startTime)/1000000)+"ms total run time.");
        writer.write("render.txt",result);
    }
}
