package dbModel

import (
	"github.com/speedianet/sfm/src/domain/dto"
	"github.com/speedianet/sfm/src/domain/entity"
	"github.com/speedianet/sfm/src/domain/valueObject"
)

type ResourceProfile struct {
	ID                     uint   `gorm:"primarykey"`
	Name                   string `gorm:"not null"`
	BaseSpecs              string `gorm:"not null"`
	MaxSpecs               *string
	ScalingPolicy          *string
	ScalingThreshold       *uint64
	ScalingMaxDurationSecs *uint64
	ScalingIntervalSecs    *uint64
	HostMinCapacityPercent *float64
}

func (ResourceProfile) TableName() string {
	return "resource_profiles"
}

func (model ResourceProfile) DefaultEntry() ResourceProfile {
	defaultEntity := entity.DefaultResourceProfile()
	defaultModel, _ := model.ToModel(defaultEntity)
	return defaultModel
}

func (ResourceProfile) ToModel(
	entity entity.ResourceProfile,
) (ResourceProfile, error) {
	var maxSpecsPtr *string
	if entity.MaxSpecs != nil {
		maxSpecs := entity.MaxSpecs.String()
		maxSpecsPtr = &maxSpecs
	}

	var scalingPolicyPtr *string
	if entity.ScalingPolicy != nil {
		scalingPolicy := entity.ScalingPolicy.String()
		scalingPolicyPtr = &scalingPolicy
	}

	var scalingThresholdPtr *uint64
	if entity.ScalingThreshold != nil {
		scalingThreshold := uint64(*entity.ScalingThreshold)
		scalingThresholdPtr = &scalingThreshold
	}

	var scalingMaxDurationSecsPtr *uint64
	if entity.ScalingMaxDurationSecs != nil {
		scalingMaxDurationSecs := uint64(*entity.ScalingMaxDurationSecs)
		scalingMaxDurationSecsPtr = &scalingMaxDurationSecs
	}

	var scalingIntervalSecsPtr *uint64
	if entity.ScalingIntervalSecs != nil {
		scalingIntervalSecs := uint64(*entity.ScalingIntervalSecs)
		scalingIntervalSecsPtr = &scalingIntervalSecs
	}

	var hostMinCapacityPercentPtr *float64
	if entity.HostMinCapacityPercent != nil {
		hostMinCapacityPercent := float64(*entity.HostMinCapacityPercent)
		hostMinCapacityPercentPtr = &hostMinCapacityPercent
	}

	return ResourceProfile{
		ID:                     uint(entity.Id.Get()),
		Name:                   entity.Name.String(),
		BaseSpecs:              entity.BaseSpecs.String(),
		MaxSpecs:               maxSpecsPtr,
		ScalingPolicy:          scalingPolicyPtr,
		ScalingThreshold:       scalingThresholdPtr,
		ScalingMaxDurationSecs: scalingMaxDurationSecsPtr,
		ScalingIntervalSecs:    scalingIntervalSecsPtr,
		HostMinCapacityPercent: hostMinCapacityPercentPtr,
	}, nil
}

func (model ResourceProfile) ToEntity() (entity.ResourceProfile, error) {
	rpId, err := valueObject.NewResourceProfileId(model.ID)
	if err != nil {
		return entity.ResourceProfile{}, err
	}

	name, err := valueObject.NewResourceProfileName(model.Name)
	if err != nil {
		return entity.ResourceProfile{}, err
	}

	baseSpecs, err := valueObject.NewContainerSpecsFromString(model.BaseSpecs)
	if err != nil {
		return entity.ResourceProfile{}, err
	}

	var maxSpecsPtr *valueObject.ContainerSpecs
	if model.MaxSpecs != nil {
		maxSpecs, err := valueObject.NewContainerSpecsFromString(*model.MaxSpecs)
		if err != nil {
			return entity.ResourceProfile{}, err
		}
		maxSpecsPtr = &maxSpecs
	}

	var scalingPolicyPtr *valueObject.ScalingPolicy
	if model.ScalingPolicy != nil {
		scalingPolicy, err := valueObject.NewScalingPolicy(*model.ScalingPolicy)
		if err != nil {
			return entity.ResourceProfile{}, err
		}
		scalingPolicyPtr = &scalingPolicy
	}

	var scalingThresholdPtr *uint64
	if model.ScalingThreshold != nil {
		scalingThreshold := uint64(*model.ScalingThreshold)
		scalingThresholdPtr = &scalingThreshold
	}

	var scalingMaxDurationSecsPtr *uint64
	if model.ScalingMaxDurationSecs != nil {
		scalingMaxDurationSecs := uint64(*model.ScalingMaxDurationSecs)
		scalingMaxDurationSecsPtr = &scalingMaxDurationSecs
	}

	var scalingIntervalSecsPtr *uint64
	if model.ScalingIntervalSecs != nil {
		scalingIntervalSecs := uint64(*model.ScalingIntervalSecs)
		scalingIntervalSecsPtr = &scalingIntervalSecs
	}

	var hostMinCapacityPercentPtr *valueObject.HostMinCapacity
	if model.HostMinCapacityPercent != nil {
		hostMinCapacityPercent, err := valueObject.NewHostMinCapacity(
			*model.HostMinCapacityPercent,
		)
		if err != nil {
			return entity.ResourceProfile{}, err
		}
		hostMinCapacityPercentPtr = &hostMinCapacityPercent
	}

	return entity.NewResourceProfile(
		rpId,
		name,
		baseSpecs,
		maxSpecsPtr,
		scalingPolicyPtr,
		scalingThresholdPtr,
		scalingMaxDurationSecsPtr,
		scalingIntervalSecsPtr,
		hostMinCapacityPercentPtr,
	)
}

func (ResourceProfile) FromAddDtoToModel(
	dto dto.AddResourceProfile,
) (ResourceProfile, error) {
	var maxSpecsPtr *string
	if dto.MaxSpecs != nil {
		maxSpecs := dto.MaxSpecs.String()
		maxSpecsPtr = &maxSpecs
	}

	var scalingPolicyPtr *string
	if dto.ScalingPolicy != nil {
		scalingPolicy := dto.ScalingPolicy.String()
		scalingPolicyPtr = &scalingPolicy
	}

	var scalingThresholdPtr *uint64
	if dto.ScalingThreshold != nil {
		scalingThreshold := uint64(*dto.ScalingThreshold)
		scalingThresholdPtr = &scalingThreshold
	}

	var scalingMaxDurationSecsPtr *uint64
	if dto.ScalingMaxDurationSecs != nil {
		scalingMaxDurationSecs := uint64(*dto.ScalingMaxDurationSecs)
		scalingMaxDurationSecsPtr = &scalingMaxDurationSecs
	}

	var scalingIntervalSecsPtr *uint64
	if dto.ScalingIntervalSecs != nil {
		scalingIntervalSecs := uint64(*dto.ScalingIntervalSecs)
		scalingIntervalSecsPtr = &scalingIntervalSecs
	}

	var hostMinCapacityPercentPtr *float64
	if dto.HostMinCapacityPercent != nil {
		hostMinCapacityPercent := float64(*dto.HostMinCapacityPercent)
		hostMinCapacityPercentPtr = &hostMinCapacityPercent
	}

	return ResourceProfile{
		Name:                   dto.Name.String(),
		BaseSpecs:              dto.BaseSpecs.String(),
		MaxSpecs:               maxSpecsPtr,
		ScalingPolicy:          scalingPolicyPtr,
		ScalingThreshold:       scalingThresholdPtr,
		ScalingMaxDurationSecs: scalingMaxDurationSecsPtr,
		ScalingIntervalSecs:    scalingIntervalSecsPtr,
		HostMinCapacityPercent: hostMinCapacityPercentPtr,
	}, nil
}
