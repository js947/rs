# RESCALE CLI DEMO

This tool implements many of the features you might need in order to setup and run jobs on the Rescale ScaleX platform.

## Analysis types

To identify the software required to run an analysis we can use the `analysis` command. This can be used either to list all analysis types

```
$ rs analysis
[... all the 100s of packages supported by rescale ...]
```

or to search for a particular piece of software

```
$ rs analysis openfoam
278 analyses indexed
5 matches
  0             openfoam
OpenFOAM
<p><b>OpenFOAM</B> is the leading free, open source software for CFD, owned by
the <B>OpenFOAM Foundation</b> and distributed exclusively under the General
Public Licence (GPL). The GPL gives users the freedom to modify and redistribute
the software and a guarantee of continued free use, within the terms of the
licence.</p>

  1          foam_extend
foam-extend
<p>The <b>foam-extend</b> project is a fork of the OpenFOAM open source library
for Computational Fluid Dynamics (CFD).<br><b>OpenFOAM</b> is a C++ toolbox for
the development of customized numerical solvers, and pre-/post-processing
utilities for the solution of continuum mechanics problems, including
computational fluid dynamics (CFD).</p>

  2        openfoam_plus
OpenFOAM+
<p><b>OpenFOAM</b> is the free, open source CFD software released and developed
primarily by <b>OpenCFD</b> Ltd since 2004. It has a large user base across most
areas of engineering and science, from both commercial and academic
organisations. OpenFOAM has an extensive range of features to solve anything
from complex fluid flows involving chemical reactions, turbulence and heat
transfer, to acoustics, solid mechanics and electromagnetics.</p>

[... other custom versions of openfoam ...]
```