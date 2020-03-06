# RESCALE CLI DEMO

This tool implements many of the features you might need in order to setup and run jobs on the Rescale ScaleX platform.

## Example use cases

We present two example use cases both of which demonstrate a simple, single-step openfoam job

- [Airfoil2D](https://github.com/js947/rs/tree/master/examples/airfoil2D)
- [MotorBike](https://github.com/js947/rs/tree/master/examples/motorbike)

However, accessing the API makes many more things possbile; for example: 

- compiling a custom app and then using the compiled binary in other jobs
- multi-stage jobs; for example running a simluation and then post-processing the data at scale

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
  output. We can also stream the logs directly to the local termial.

5. Collect outputs

  When the job completes we may wish to download the output files for futher analysis.

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

To query the available versions of an analysis, we can use the `versions` command with the analysis *code*. The *code* is the simple ID - i.e. openfoam rather than openFOAM, openfoam_plus rather than OpenFOAM+. 

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

It is necessary to pick a core type that has support for the specific application that you are trying to run, and you can filter with the application and version code to see the core types that support that application version

```
$ rs cores openfoam 7-intelmpi
[... lots of core types ...]
```

### Job submission

```
$ rs submit
[... files uploaded ...]
job create <jobid>
job submit <jobid>
```

### Job files

```
$ rs files list --jobid=<jobid>
[... list output files for job ...]
```

## Further Information

The full capabilities of each command is described in the wiki pages:

- Files `$ rs file [upload,list,delete]` [](https://github.com/js947/rs.wiki/file)
- Jobs `$ rs job [list,rename,delete]` [](https://github.com/js947/rs.wiki/job)
