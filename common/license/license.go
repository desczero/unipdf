//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package license helps manage commercial licenses and check if they are valid for the version of unipdf used.
package license ;import _a "github.com/unidoc/unipdf/v3/internal/license";

// GetMeteredState checks the currently used metered document usage status,
// documents used and credits available.
func GetMeteredState ()(_a .MeteredStatus ,error ){return _a .GetMeteredState ()};

// SetLicenseKey sets and validates the license key.
func SetLicenseKey (content string ,customerName string )error {return _a .SetLicenseKey (content ,customerName );};

// SetMeteredKey sets the metered API key required for SaaS operation.
// Document usage is reported periodically for the product to function correctly.
func SetMeteredKey (apiKey string )error {return _a .SetMeteredKey (apiKey )};

// MakeUnlicensedKey returns a default key.
func MakeUnlicensedKey ()*LicenseKey {return _a .MakeUnlicensedKey ()};

// GetLicenseKey returns the currently loaded license key.
func GetLicenseKey ()*LicenseKey {return _a .GetLicenseKey ()};const (LicenseTierUnlicensed =_a .LicenseTierUnlicensed ;LicenseTierCommunity =_a .LicenseTierCommunity ;LicenseTierIndividual =_a .LicenseTierIndividual ;LicenseTierBusiness =_a .LicenseTierBusiness ;
);

// SetMeteredKeyPersistentCache sets the metered License API Key persistent cache.
// Default value 'true', set to `false` will report the usage immediately to license server,
// this can be used when there's no access to persistent data storage.
func SetMeteredKeyPersistentCache (val bool ){_a .SetMeteredKeyPersistentCache (val )};

// LicenseKey represents a loaded license key.
type LicenseKey =_a .LicenseKey ;