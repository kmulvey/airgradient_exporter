#!/bin/sh
systemctl daemon-reload
systemctl enable airgradient_exporter.service
systemctl start airgradient_exporter.service