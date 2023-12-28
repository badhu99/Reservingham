import React, { useRef, useEffect } from "react";
import { ElementCircle, ElementRect, CanvasElement } from "../../interfaces/shapes";

interface ICanvasProps {
  Elements: CanvasElement[];
  UpdateElements: (updatedElements: CanvasElement[]) => void;
  SetClickedElement : (element: CanvasElement | undefined) => void;
}

const Canvas: React.FC<ICanvasProps> = ({ Elements, UpdateElements, SetClickedElement}) => {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  var ctx = canvasRef.current?.getContext("2d");
  var offsetX = canvasRef.current?.getBoundingClientRect()?.left || 0;
  var offsetY = canvasRef.current?.getBoundingClientRect()?.top || 0;
  var canvasWidth = canvasRef.current?.width || 0;
  var canvasHeight = canvasRef.current?.height || 0;

  // drag related variables
  const [dragok, setDragok] = React.useState(false);
  const [startX, setStartX] = React.useState(0);
  const [startY, setStartY] = React.useState(0);

  const [hasMoved, setHasMoved] = React.useState(false);



  const draw = () => {
    clear();
    for (let i = 0; i < Elements.length; i++) {
      if ("width" in Elements[i]) {
        rect(Elements[i] as ElementRect);
      } 
      else {
        circle(Elements[i] as ElementCircle);
      }
    }
  };


  const rect = (r: ElementRect) => {
    ctx!.fillStyle = r.fill;
    ctx?.fillRect(r.x, r.y, r.width, r.height);
  };

  const circle = (c: ElementCircle) => {
    ctx!.fillStyle = c.fill;
    ctx?.beginPath();
    ctx?.arc(c.x, c.y, c.r, 0, Math.PI * 2);
    ctx?.closePath();
    ctx?.fill();
  };

  const clear = () => {
    ctx?.clearRect(0, 0, canvasWidth, canvasHeight);
  };

  // useEffect to handle canvas setup
  useEffect(() => {
    if (ctx) {
      draw();
    }
    else if(ctx === undefined) {
      ctx = canvasRef.current?.getContext("2d");
      offsetX = canvasRef.current?.getBoundingClientRect()?.left || 0;
      offsetY = canvasRef.current?.getBoundingClientRect()?.top || 0;
      canvasWidth = canvasRef.current?.width || 0;
      canvasHeight = canvasRef.current?.height || 0;
    }
  }, [ctx, Elements, canvasWidth, canvasHeight]);

  // event handlers
  const myDown = (e: React.MouseEvent) => {
    const {mx, my} = getClickPosition(e);
    setDragok(false);
    for (let i = 0; i < Elements.length; i++) {
      const s = Elements[i];
      if ("width" in s) {
        if (dragok === false && mx > s.x && mx < s.x + s.width && my > s.y && my < s.y + s.height) {
          setDragok(true);
          s.isDragging = true;
          setHasMoved(false);
        }
      } else {
        const dx = s.x - mx;
        const dy = s.y - my;
        if (dragok === false && dx * dx + dy * dy < s.r * s.r) {
          setDragok(true);
          s.isDragging = true;
          setHasMoved(false);
        }
      }
    }

    setStartX(mx);
    setStartY(my);
  };

  const myUp = (e: React.MouseEvent) => {
    getClickPosition(e);

    setDragok(false);
    let element: CanvasElement | undefined = undefined;
    for (let i = 0; i < Elements.length; i++) {
      if(Elements[i].isDragging === true) {
        element = Elements[i];
      }
      Elements[i].isDragging = false;
    }

    if(hasMoved) {
      SetClickedElement(undefined);
    }
    else{
      SetClickedElement(element)
    }
    setHasMoved(false);
  };

  const myMove = (e: React.MouseEvent) => {
    if (dragok === true) {
      const {mx, my} = getClickPosition(e);

      const dx = mx - startX;
      const dy = my - startY;

      for (let i = 0; i < Elements.length; i++) {
        const s = Elements[i];
        if (s.isDragging === true) {
          s.x += dx;
          s.y += dy;
          setHasMoved(true);
        }
      }

      setStartX(mx);
      setStartY(my);
      UpdateElements(Elements);
    }
  }

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
