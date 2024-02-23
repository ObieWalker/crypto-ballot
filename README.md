
<a name="readme-top"></a>

[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">An application that simulates voting using Blockchain.</h3>

</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This project displays features and implementation of a blockchain voting application. A user can simulate the process and see the mining process alter the difficulty to improve the Proof Of Work and broadcast blocks to other connected nodes that will then update the blockchain so all nodes have correlating data on their blockchains.



<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

This application was built primarily with Golang as it is the backend application of a supposed voting app. The endpoints can be tested with postman and data can be gotten using your browser.

* [![Golang][golang]][Go-url]
* [![Postman][postman]][Postman-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Running this project is fairly straightforward. 

To get a local copy up and running follow these simple example steps.


### Installation


1. Install Golang on your system
2. Clone the repo
   ```sh
   git clone https://github.com/ObieWalker/crypto-ballot.git
   ```
3. Install Go packages
   ```sh
   go mod download
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage

Once full installation is complete, open a terminal navigate to the `cmd` folder and run `go run .go`. The application should start running.

Open another terminal (more than one if possible) and navigate to the `nodes` folder and run `go run .go`. 

You can then open postman and use the endpoint `localhost:8081`. As every new vote is created in postman, you will see it broadcast to all connected nodes in real time simulating the blockchain process. 

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## API Endpoints

***Create Polling Unit***
* `POST` localhost:8081/pollingunit
JSON Body
```
{
    "PollingUnit": {sample_polling_unit}
}
```

***Create Vote***
* `POST` localhost:8081
JSON Body
```
{
    "PollingUnit": "373",
    "VoterId": "3233",
    "ElectionType": "presidential",
    "Selection": "2"
}
```

***Get Polling Unit Votes***
* `GET` localhost:8081/puvotes/{sample_polling_unit}


***Get Polling Unit Blockchain***
* `GET` localhost:8081/{sample_polling_unit}


***Get Total Votes***
* `GET` localhost:8081/votes


***Total Votes By Polling Unit***
* `GET` localhost:8081/puvotes


***Get all blockchains***
* `GET` localhost:8081/


***Validate Polling Unit Chain***
* `POST` localhost:8081/validate/{sample_polling_unit}



<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Obi Walker - [My LinkedIn](https://www.linkedin.com/in/obi-walker-a9b94371/) - obinnawalker@gmail.com


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [Choose an Open Source License](https://choosealicense.com)
* [Best-README-Template](https://github.com/othneildrew/Best-README-Template)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[linkedin-url]: https://www.linkedin.com/in/obi-walker-a9b94371/
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/ObieWalker/crypto-ballot/blob/develop/LICENSE.txt
[Golang]: https://img.shields.io/badge/go-000000?style=for-the-badge&logo=go&logoColor=##FF6C37
[Go-url]: https://go.dev/
[postman]: https://img.shields.io/badge/postman-000000?style=for-the-badge&logo=go&logoColor=#007D9C
[Postman-url]: https://www.postman.com/