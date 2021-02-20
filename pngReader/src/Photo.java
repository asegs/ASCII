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
