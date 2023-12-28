export interface RectShape {
    id: string;
    name:string;
    x: number;
    y: number;
    width: number;
    height: number;
    fill: string;
    isDragging: boolean;
  }
  
  export interface CircleShape {
    id: string;
    name:string;
    x: number;
    y: number;
    r: number;
    fill: string;
    isDragging: boolean;
  }

  export type Shape = RectShape | CircleShape;