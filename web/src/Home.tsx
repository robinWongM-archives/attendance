import React from "react";
import PunchInCard from "./PunchInCard";

function Home() {
  return (
    <div className="flex-1 bg-white rounded-l-3xl p-12 text-4xl">
      <h1>👋 签个到吧，贾同学</h1>
      <PunchInCard />
    </div>
  );
}

export default Home;
