import React from "react";
import { Box, Button, Container, Grid } from "@awsui/components-react";
import style from "./Goods.module.scss";
import { url } from "inspector";
import img_model_asr from '../images/asr.png';
import img_model_mt from '../images/machine_translation.png';
import img_model_ic from '../images/img_classification.png';
import img_model_tsp from '../images/time_series_prediction.png';
import img_data_imagenet from '../images/imagenet.png';
import img_data_europarl from '../images/europarl.png';
import img_data_celeba from '../images/celeba.png';
import img_data_mozilla from '../images/common_voice.png';
import img_power_gpulab from '../images/gpulab.png';
import img_power_leadergpu from '../images/leadergpu.png';

const img_style = {
	width: '200px', 
	height: '200px',
	}
const Goods = () => {
  const model_list = [
    {
      title: "Automatic Speech Recognition Model",
      body: "This model is an automatic speech recognition system implemented in PyTorch, supporting both English and German languages.",
      url: "../asr.png",
      Image: img_model_asr,
    },
    {
      title: "Machine Translation Model",
      body: "This model is a Transformer-based machine translation model implemented in TensorFlow.",
      url: "/path/to/the/product",
      Image: img_model_mt,
    },
    {
      title: "Image Classification Model",
      body: "This model is a neural network implemented in PyTorch for common animal recognition.",
      url: "/path/to/the/product",
      Image: img_model_ic,
    },
    {
      title: "Time Series Prediction Model",
      body: "This service uses Prophet and Statsmodel to forecast and analyze points of a given time series. ",
      url: "/path/to/the/product",
      Image: img_model_tsp,
    },
  ];


  const data_list = [
    {
      title: "ImageNet Dataset",
      body: "ImageNet is an image database organized according to the WordNet hierarchy (currently only the nouns), in which each node of the hierarchy is depicted by hundreds and thousands of images. The project has been instrumental in advancing computer vision and deep learning research. The data is available for free to researchers for non-commercial use.",
      url: "https://www.image-net.org/download.php",
      Image: img_data_imagenet,
    },
    {
      title: "Large-scale CelebFaces Attributes (CelebA) Dataset",
      body: "CelebFaces Attributes Dataset (CelebA) is a large-scale face attributes dataset with more than 200K celebrity images, each with 40 attribute annotations. The images in this dataset cover large pose variations and background clutter. CelebA has large diversities, large quantities, and rich annotations",
      url: "https://mmlab.ie.cuhk.edu.hk/projects/CelebA.html",
      Image: img_data_celeba,
    },
    {
      title: "European Parliament Proceedings Parallel Corpus",
      body: "The Europarl parallel corpus is extracted from the proceedings of the European Parliament. It includes versions in 21 European languages: Romanic (French, Italian, Spanish, Portuguese, Romanian), Germanic (English, Dutch, German, Danish, Swedish), Slavik (Bulgarian, Czech, Polish, Slovak, Slovene), Finni-Ugric (Finnish, Hungarian, Estonian), Baltic (Latvian, Lithuanian), and Greek.",
      url: "https://www.statmt.org/europarl/",
      Image: img_data_europarl,
    },
    {
      title: "Mozilla Common Voice",
      body: "Mozilla claims to have the largest human speech dataset available, with a current dataset of 29 different languages, including Chinese, and nearly 2454 hours (1965 of which are verified) of recorded speech data collected from more than 40,000 contributors. And we have made an open commitment to expose our collection of high-quality speech data to startups, researchers, and anyone interested in speech technology.",
      url: "https://commonvoice.mozilla.org/zh-CN",
      Image: img_data_mozilla,
    },
  ];


  const power_list = [
    {
      title: "GPULab",
      body: "GPULab is a turnkey JupyterLab Notebook environment atop a feature-packed Ubuntu Linux operating system. GPULab provides data scientists and research teams a dedicated Nvidia K80 GPU for a flat fee.",
      url: "https://gpulab.io/",
      Image: img_power_gpulab,
    },
    {
      title: "LeaderGPU",
      body: "LeaderGPU®, a LeaderTelecom initiative, focuses on providing GPU-power for high-performance computing, machine learning, rendering, etc. We provide bare-metal servers, not VPS. The servers are optimised for software that is required for machine learning and training of neural networks (Tensorflow™, PyTorch, etc.). In addition to this, users can install their own software.",
      url: "https://www.leadergpu.com/",
      Image: img_power_leadergpu,
    },
    // {
    //   title: "Image Classification Model",
    //   body: "description",
    //   url: "/path/to/the/product",
    //   Image: "images/image_classification.png",
    // },
    // {
    //   title: "Time Series Prediction Model",
    //   body: "description",
    //   url: "/path/to/the/product",
    //   Image: "images/time_series_prediction.png",
    // },
  ];


  return (
    <>
      <Box variant="h1" margin={{ top: "xxl" }}>
        Available Models
      </Box>
      <Grid
        gridDefinition={[
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
        ]}
      >
        {model_list.map((good, index) => (
          <Container key={index} className={style["quick-box"]}>
            <Box variant="h3" padding={{ top: "n" }} fontWeight="bold">
              {good.title}
            </Box>
            <img src={good.Image}  alt={good.title} style = {img_style}/>
            <Box variant="p">{good.body}</Box>
            <Button href={good.url} variant="normal">
              learn more
            </Button>
          </Container>
        ))}
      </Grid>

      <Box variant="h1" margin={{ top: "xxl" }}>
        Available Datasets
      </Box>
      <Grid
        gridDefinition={[
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
        ]}
      >
        {data_list.map((good, index) => (
          <Container key={index} className={style["quick-box"]}>
            <Box variant="h3" padding={{ top: "n" }} fontWeight="bold">
              {good.title}
            </Box>
            <img src={good.Image} alt={good.title} style = {img_style}/>
            <Box variant="p">{good.body}</Box>
            <Button href={good.url} variant="normal">
              learn more
            </Button>
          </Container>
        ))}
      </Grid>

      <Box variant="h1" margin={{ top: "xxl" }}>
        Available Compute Power
      </Box>
      <Grid
        gridDefinition={[
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
        ]}
      >
        {power_list.map((good, index) => (
          <Container key={index} className={style["quick-box"]}>
            <Box variant="h3" padding={{ top: "n" }} fontWeight="bold">
              {good.title}
            </Box>
            <img src={good.Image} alt={good.title} style = {img_style}/>
            <Box variant="p">{good.body}</Box>
            <Button href={good.url} variant="normal">
              learn more
            </Button>
          </Container>
        ))}
      </Grid>

    </>
    
  );
};

export default Goods;
