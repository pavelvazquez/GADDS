# Global Accessible Distributed Data Sharing Platform

Browser-based client interface to simplify the implementation of FAIRness in the data collection and storage process in life sciences.

## Getting Started

This a demo version to be instatiated on single hardware. 

Usage: 

  start.sh -h|--help print help

    -m - one of 'up', 'down' or 'restart'

      - 'up' - bring up the demo platform

      - 'down' - clear the demo platform
      
      - 'restart' - restart the demo platform

Example, instatiate the platform: 

    ./start.sh -m up

*Note that when turning down the platform, the volumes are not removed. To do so type: "docker volume prune". Warning: This will erase the blockchain ledger!



### Prerequisites
*Tested on Unix or Mac OS systems only

-Docker Engine 19.03.13 or higher

-Docker-compose 1.27.4 or higher

## Authors

* **Pavel Vazquez Faci** - [Personal](https://www.med.uio.no/hth/english/people/postdocs/pavelva/index.html)

See also the list of [contributors](CONTRIBUTORS.md) who participated in this project.

## License

This project is licensed under the Apache 2.0 License- see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Project founded by the Hybrid Technology Hub - Centre for Organ on a Chip-Technology
* This project uses modified code from https://github.com/olegabu/fabric-starter, licensed under the Apache 2.0 license. 
