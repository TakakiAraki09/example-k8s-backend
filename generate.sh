#!/bin/bash
atlas schema inspect --env local --format '{{ sql . }}' > schema/db/schema.sql
