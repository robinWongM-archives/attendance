import React from "react";
export function Button({
  className,
  Icon,
  Content,
}: {
  className?: string;
  Icon: string;
  Content: string;
}) {
  return (
    <button
      className={
        "text-blue-500 text-lg appearance-none focus:outline-none bg-white text-center pl-4 pr-5 py-2 rounded-full " +
        className
      }
    >
      <img className="w-8 h-8 inline-block align-middle" src={Icon}></img>{" "}
      <span className="align-middle">{Content}</span>
    </button>
  );
}
