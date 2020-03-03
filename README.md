# otdd-test-runner

The otdd test runner runs in your development node.

It fetches online-recorded test cases from otddserver and run them locally, then reports results back to otddserver.
When a test is run, the inbound request is send by testrunner to local port, the outbound calls are then mocked by the test-runner using data fetched from otddserver. 
Thus the development code is run as if "online" using real online data, dramatically increase the developing and testing efficiency.

The otdd test runner run in two modes:

- *docker mode*. The otdd test runner is run within a docker, and your development docker share its network with the runner docker. They acts like the istio sidecare mode.
- *plain mode*. The otdd runner is run within a seperate otdd-test-runner user in your development box.

Please refer to [get started](https://otdd.io/getstarted) to learn about the setup.


