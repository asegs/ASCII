import java.io.IOException;

public class ToAscii {
    public char photoToAscii(Photo photo){
        if (photo==null){
            return ' ';
        }
        double[][][] imageData = photo.getImgData();
        int boxCount = imageData.length*imageData[0].length;
        double totalColorNum = 0.0;
        for (int row=0;row<imageData.length;row++){
            for (int col=0;col<imageData[0].length;col++){
                for (int i=0;i<3;i++){
                    totalColorNum+=imageData[row][col][i];
                }
            }
        }
        double avgDarkness = totalColorNum/(boxCount*3);
        if (avgDarkness<=30){
            return 'M';
        }
        else if(avgDarkness<=45){
            return 'N';
        }
        else if(avgDarkness<=60){
            return 'm';
        }
        else if(avgDarkness<=80){
            return 'd';
        }
        else if(avgDarkness<=100){
            return 'y';
        }
        else if (avgDarkness<=120){
            return 's';
        }
        else if (avgDarkness<=140){
            return '+';
        }
        else if (avgDarkness<=160){
            return ':';
        }
        else if (avgDarkness<=180){
            return '/';
        }
        else if (avgDarkness<=200){
            return '-';
        }
        else if (avgDarkness<=220){
            return '`';
        }
        else{
            return ' ';
        }
    }
}
