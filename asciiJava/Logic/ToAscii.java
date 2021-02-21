package com.example.ascii.Logic;

public class ToAscii {
    public char photoToAscii(Photo photo,boolean inverse){
        if (photo==null){
            return ' ';
        }
        char[] returns = new char[]{'M','N','m','/','d','y','s','+',':','-','`',' '};
        double[][][] imageData = photo.getImgData();
        int boxCount = imageData.length*imageData[0].length;
        double totalColorNum = 0.0;
        int toSubtract = inverse ? 0 : 11;
        for (int row=0;row<imageData.length;row++){
            for (int col=0;col<imageData[0].length;col++){
                for (int i=0;i<3;i++){
                    try {
                        totalColorNum+=imageData[row][col][i];
                    }
                    catch (NullPointerException npe){
                        return ' ';
                    }

                }
            }
        }
        double avgDarkness = totalColorNum/(boxCount*3);
        if (avgDarkness<=30){
            return returns[Math.abs(toSubtract-11)];
        }
        else if(avgDarkness<=45){
            return returns[Math.abs(toSubtract-10)];
        }
        else if(avgDarkness<=60){
            return returns[Math.abs(toSubtract-9)];
        }
        else if(avgDarkness<=80){
            return returns[Math.abs(toSubtract-8)];
        }
        else if(avgDarkness<=100){
            return returns[Math.abs(toSubtract-7)];
        }
        else if (avgDarkness<=120){
            return returns[Math.abs(toSubtract-6)];
        }
        else if (avgDarkness<=140){
            return returns[Math.abs(toSubtract-5)];
        }
        else if (avgDarkness<=160){
            return returns[Math.abs(toSubtract-4)];
        }
        else if (avgDarkness<=180){
            return returns[Math.abs(toSubtract-3)];
        }
        else if (avgDarkness<=200){
            return returns[Math.abs(toSubtract-2)];
        }
        else if (avgDarkness<=220){
            return returns[Math.abs(toSubtract-1)];
        }
        else{
            return returns[Math.abs(toSubtract)];
        }
    }
}
