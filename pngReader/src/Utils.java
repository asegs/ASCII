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
}
