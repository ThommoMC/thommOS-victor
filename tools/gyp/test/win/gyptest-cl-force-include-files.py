#!/usr/bin/env python2

# Copyright (c) 2012 Google Inc. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""
Make sure ForcedIncludeFiles option is functional.
"""

import TestGyp

import sys

if sys.platform == 'win32':
  test = TestGyp.TestGyp(formats=['msvs', 'ninja'])

  CHDIR = 'compiler-flags'
  test.run_gyp('force-include-files.gyp', chdir=CHDIR)
  test.build('force-include-files.gyp', test.ALL, chdir=CHDIR)

  test.pass_test()
