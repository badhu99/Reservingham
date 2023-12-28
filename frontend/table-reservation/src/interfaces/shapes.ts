export interface ElementRect {
    id: string;
    name:string;
    x: number;
    y: number;
    width: number;
    height: number;
    fill: string;
    isDragging: boolean;
  }
  
  export interface ElementCircle {
    id: string;
    name:string;
    x: number;
    y: number;
    r: number;
    fill: string;
    isDragging: boolean;
  }

  export type CanvasElement = ElementRect | ElementCircle;