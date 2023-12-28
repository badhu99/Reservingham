import React, { useRef, useEffect } from "react";
import { CircleShape, RectShape, Shape } from "../../interfaces/shapes";

interface ICanvasProps {
  shapes: Shape[];
  updateShapes: (updatedShapes: Shape[]) => void;
  updateShape: (updatedShape: Shape) => void;
  setShapes: React.Dispatch<React.SetStateAction<Shape[]>>;
  onShapesChange: (newShapes: Shape[]) => void;
}

const Canvas: React.FC<ICanvasProps> = ({ shapes, updateShapes, updateShape, setShapes, onShapesChange}) => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
const shapesRef = useRef<Shape[]>(shapes);


  // get canvas related references
  var canvas = canvasRef.current;
  var ctx = canvas?.getContext("2d");
  var BB = canvas?.getBoundingClientRect();
  var offsetX = BB?.left || 0;
  var offsetY = BB?.top || 0;
  var WIDTH = canvas?.width || 0;
  var HEIGHT = canvas?.height || 0;

  // drag related variables
  let dragok = false;
  let startX: number;
  let startY: number;

  const draw = () => {
    clear();
    for (let i = 0; i < shapes.length; i++) {
      if ("width" in shapes[i]) {
        rect(shapes[i] as RectShape);
      } 
      else {
        circle(shapes[i] as CircleShape);
      }
    }
  };

  const rect = (r: RectShape) => {
    ctx!.fillStyle = r.fill;
    ctx?.fillRect(r.x, r.y, r.width, r.height);
  };

  const circle = (c: CircleShape) => {
    ctx!.fillStyle = c.fill;
    ctx?.beginPath();
    ctx?.arc(c.x, c.y, c.r, 0, Math.PI * 2);
    ctx?.closePath();
    ctx?.fill();
  };

  const clear = () => {
    ctx?.clearRect(0, 0, WIDTH, HEIGHT);
  };

  // useEffect to handle canvas setup
  useEffect(() => {
    if (ctx) {
      console.log("ctx is not undefined");
      draw();
    }
    else if(ctx === undefined) {
      console.log("ctx is undefined");
      canvas = canvasRef.current;
      ctx = canvas?.getContext("2d");
      BB = canvas?.getBoundingClientRect();
      offsetX = BB?.left || 0;
      offsetY = BB?.top || 0;
      WIDTH = canvas?.width || 0;
      HEIGHT = canvas?.height || 0;
      // draw();
    }
  }, [ctx, WIDTH, HEIGHT]);

  // event handlers
  const myDown = (e: React.MouseEvent) => {
    const {mx, my} = getClickPosition(e);
    dragok = false;
    for (let i = 0; i < shapes.length; i++) {
      const s = shapes[i];
      if ("width" in s) {
        if (dragok === false && mx > s.x && mx < s.x + s.width && my > s.y && my < s.y + s.height) {
          dragok = true;
          s.isDragging = true;
        }
      } else {
        const dx = s.x - mx;
        const dy = s.y - my;
        if (!dragok && dx * dx + dy * dy < s.r * s.r) {
          dragok = true;
          s.isDragging = true;
        }
      }
    }

    startX = mx;
    startY = my;
  };

  const myUp = (e: React.MouseEvent) => {
    getClickPosition(e);

    dragok = false;
    for (let i = 0; i < shapes.length; i++) {
      shapes[i].isDragging = false;
    }

    // updateShapes(shapes);
  };

  const myMove = (e: React.MouseEvent) => {
    if (dragok) {
      const {mx, my} = getClickPosition(e);

      const dx = mx - startX;
      const dy = my - startY;

      for (let i = 0; i < shapes.length; i++) {
        const s = shapes[i];
        if (s.isDragging === true) {
          s.x += dx;
          s.y += dy;
          // console.log("sx, sy: ", s.x, s.y);
          console.log("s", s);
          // updateShapeCoordinates(i, s.x, s.y);
        }
      }

      draw();
      startX = mx;
      startY = my;
      // updateShapes(shapes);
    }
  }


const updateShapeCoordinates = (shapeIndex: number, newX: number, newY: number) => {
  const updatedShapes = [...shapesRef.current];
  const shapeToUpdate = updatedShapes[shapeIndex];
  shapeToUpdate.x = newX;
  shapeToUpdate.y = newY;
  shapesRef.current = updatedShapes;
  // onShapesChange(updatedShapes);
  updateShapes(updatedShapes);
};

  // const onClick = (e: React.MouseEvent) => {

  //   const {mx, my} = getClickPosition(e);

  //   for (let i = 0; i < shapes.length; i++) {
  //     const s = shapes[i];
  //     if ("width" in s) {
  //       if (mx > s.x && mx < s.x + s.width && my > s.y && my < s.y + s.height) {
  //         console.log(s.name);
  //       }
  //     } else {
  //       const dx = s.x - mx;
  //       const dy = s.y - my;
  //       if (dx * dx + dy * dy < s.r * s.r) {
  //         console.log(s.name);
  //       }
  //     }
  //   }
  // }

  const getClickPosition = (e: React.MouseEvent) => {

    e.preventDefault();
    e.stopPropagation();

    const mx = parseInt((e.clientX - offsetX).toString());
    const my = parseInt((e.clientY - offsetY).toString());

    return { mx, my };
  }

  return (
    <canvas
      ref={canvasRef}
      id="canvas"
      onMouseDown={myDown}
      onMouseUp={myUp}
      onMouseMove={myMove}
      width={1000}
      height={700}
      style={{ border: "1px solid #000" }}
    />
  );
};

export default Canvas;
