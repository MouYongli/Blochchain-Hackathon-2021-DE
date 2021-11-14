import { FC } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import AssetPage from "./asset/AssetPage";
import WelcomePage from "./home/WelcomePage";
const Router: FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/asset" element={<AssetPage />} />
        <Route path="/" element={<WelcomePage />} />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
