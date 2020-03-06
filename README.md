# RESCALE CLI DEMO

This tool implements many of the features you might need in order to setup and run jobs on the Rescale ScaleX platform.

## Example use cases

We present two example use cases both of which demonstrate a simple, single-step openfoam job

- [Airfoil2D](https://github.com/js947/rs/tree/master/examples/airfoil2D)
- [MotorBike](https://github.com/js947/rs/tree/master/examples/motorbike)

However, accessing the API makes many more things possible; for example: 

- compiling a custom app and then using the compiled binary in other jobs
- multi-stage jobs; for example running a simulation and then post-processing the data at scale

## Walkthough

The typical workflow to design and submit a job will require the following
steps

1. Choose the analysis software required.

  Rescale supports over 200 software packages and we can use this tool to
  search on the command line for a particular package, and also list the
  available versions of each package.

2. Choose the hardware (or 'core') type to run on.

  Rescale includes a wide range of hardware types and here we have
  functionality to list all the various core types, and filter down to those that support our chosen analysis software.

3. Submit a job.

  From local input files, `rs` can upload the input files gather all the required options and submit the job.

4. Monitoring

  Once submitted we can open a job's page on the rescale platform to view its
  output. We can also stream the logs directly to the local terminal.

5. Collect outputs

  When the job completes we may wish to download the output files for further analysis.

### Analysis types

To identify the software required to run an analysis we can use the `analysis` command. This can be used either to list all analysis types

```
$ rs analyses
[... all the 100s of packages supported by rescale ...]
```

or to search for a particular piece of software

```
$ rs analyses openfoam
281 analyses indexed
5 matches
          'openfoam' OpenFOAM
<p><b>OpenFOAM</B> is the leading free, open source software for CFD, owned by
the <B>OpenFOAM Foundation</b> and distributed exclusively under the General
Public Licence (GPL). The GPL gives users the freedom to modify and redistribute
the software and a guarantee of continued free use, within the terms of the
licence.</p>

       'foam_extend' foam-extend
<p>The <b>foam-extend</b> project is a fork of the OpenFOAM open source library
for Computational Fluid Dynamics (CFD).<br><b>OpenFOAM</b> is a C++ toolbox for
the development of customized numerical solvers, and pre-/post-processing
utilities for the solution of continuum mechanics problems, including
computational fluid dynamics (CFD).</p>

     'openfoam_plus' OpenFOAM+
<p><b>OpenFOAM</b> is the free, open source CFD software released and developed
primarily by <b>OpenCFD</b> Ltd since 2004. It has a large user base across most
areas of engineering and science, from both commercial and academic
organisations. OpenFOAM has an extensive range of features to solve anything
from complex fluid flows involving chemical reactions, turbulence and heat
transfer, to acoustics, solid mechanics and electromagnetics.</p>

[... other custom versions of openfoam ...]
```

### Analysis versions

To query the available versions of an analysis, we can use the `versions`
command with the analysis *code*. The *code* is the simple ID - i.e. openfoam
rather than openFOAM, openfoam_plus rather than OpenFOAM+. 

```
$ rs analysis versions openfoam
       7 (Intel MPI)    7-intelmpi
       6 (Intel MPI)    6-intelmpi
                 5.0    5.0
                 4.1    4.1
               3.0.0    3.0.0-openmpi
               2.4.0    2.4.0-openmpi
               2.3.1    2.3.1-openmpi
               2.3.0    2.3.0-openmpi
               2.2.2    2.2.2-openmpi
               2.2.0    2.2.0-openmpi
               2.1.1    2.1.1-openmpi
```

This command returns the description along with the version code.

### Core types

To identify what kind of core (hardware) we can run on, we can use the `cores` command.

To list all of the core times available

```
$ rs cores
        code            name    description
       hpc-3            Onyx    Intel Xeon E5-2666 v3 (Haswell)
     mercury         Mercury    Intel Xeon E5-2667 v3 (Haswell)
[... total 30+ different core types ...]
```

It is necessary to pick a core type that has support for the specific
application that you are trying to run, and you can filter with the application
and version code to see the core types that support that application version

```
$ rs cores openfoam 7-intelmpi
[... lots of core types ...]
```

### Job submission

To submit a job we can `rs submit`. This will:
- pack up the jobs input files into a zip file and upload it;
- read the `rescale.yaml` file to determine the analysis software, core type and other settings; and
- submit the job.

```
$ rs submit
[... files uploaded ...]
job create <jobid>
job submit <jobid>
```

By default this will look for the job files in the current directory, and
read `rescale.yaml` from that directory. To submit a job in a different
directory we can use the `--path` option:

```
$ rs submit --path=mysimulations/motorBike
[... upload input from mysimulations/motorBike ...]
```

and to use a different `rescale.yaml` input file we can use the `--config` option:

```
$ rs submit --config=config2.yaml
[... look for input in the config2.yaml file ...]
```

The core type, number of cores, required software and other settings are
specified in the config file. For example

```
name: OpenFOAM MotorBike
core: hpc-3
numcores: 8
analysis:
- software: openfoam
  version: 7-intelmpi
  command: ./Allrun
```

where the core type and analysis software use the code rather than
human-readable name (i.e. `hpc-3` rather than `Onyx`). Some of these settings
can be overridden on the command line e.g.

```
$ rs submit --core luna --numcores 24
```

to choose a 24-core cascade lake architecture. This might allow fast
comparison of different core types and core counts to determine the most
efficient configuration for a given simulation.

### Opening a job

To investigate the progress of a job by opening it in the rescale web
interface we can use `rs job open`

```
$ rs job open <jobid>
```

with the jobid returned from the submission. This will open the web page in
the operating system's default browser.

### Job files

To investigate the output files from a given job we can use `rs job files` to
list output files, and `rs file cat` to quickly read them. For example

```
$ rs job files <jobid>
          ID            type                        name        name
      DJdFem          output          process_output.log        user/user_yXwteb/output/job_Ptczhb/run1/process_output.log
      ikLbhm          output                      U.orig        user/user_yXwteb/output/job_Ptczhb/run1/0/U.orig
      vvLbhm          output                  fixedInlet        user/user_yXwteb/output/job_Ptczhb/run1/0/include/fixedInlet
[... many other output files for job ...]
```

to get the list of output files, and

```
$ rs file cat <fileid>
[2020-03-06T12:05:46Z]: Launching ./Allrun, Working dir: /enc/uprod_Ptczhb/work/shared.  Process output follows:
[2020-03-06T12:05:47Z]: Running surfaceFeatures on /enc/uprod_Ptczhb/work/shared
[2020-03-06T12:05:59Z]: Running blockMesh on /enc/uprod_Ptczhb/work/shared
[2020-03-06T12:06:00Z]: Running decomposePar on /enc/uprod_Ptczhb/work/shared
[2020-03-06T12:06:01Z]: Running snappyHexMesh in parallel on /enc/uprod_Ptczhb/work/shared using 6 processes
[2020-03-06T12:07:18Z]: Running patchSummary in parallel on /enc/uprod_Ptczhb/work/shared using 6 processes
[2020-03-06T12:07:19Z]: Running potentialFoam in parallel on /enc/uprod_Ptczhb/work/shared using 6 processes
[2020-03-06T12:07:20Z]: Running simpleFoam in parallel on /enc/uprod_Ptczhb/work/shared using 6 processes
[2020-03-06T12:10:09Z]: Running reconstructParMesh on /enc/uprod_Ptczhb/work/shared
[2020-03-06T12:10:12Z]: Running reconstructPar on /enc/uprod_Ptczhb/work/shared
[2020-03-06T12:10:13Z]: Exited with code 0
```

to read a particular output file (here the log of an openfoam job going through its phases).

To download a larger file (perhaps a vtk output file)

```
$ rs file download <fileid>
```

will save the output file in the current directory.

## Further Information

The full capabilities of each command is described in the wiki pages:

- [Files](https://github.com/js947/rs.wiki/file) `$ rs file [upload,list,delete]`
- [Jobs](https://github.com/js947/rs.wiki/job) `$ rs job [list,rename,delete]`
